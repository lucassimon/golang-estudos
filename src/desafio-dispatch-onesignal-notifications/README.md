# One Signal Notifications pelo backend

Criar um binario que receba um payload e dispare uma request para a API do OneSignal

---

```
def notify_a_new_member(team_id, new_member_id):
    from app.challenge.models.team import ChallengeTeamUserRel, Team
    from app.users.models.user import User

    team = Team.query.get(team_id)
    new_member = User.query.get(new_member_id)
    users_to_notify = team.members.filter(
        ChallengeTeamUserRel.user_id != new_member_id
    ).all()
    users_ids = []
    kwargs = {"team_id": team.id, "challenge_id": team.challenge.id}
    notification = Notification.get_by_template_or_none("new_member")
    if not notification:
        return
    contents = notification.get_all_contents(
        new_user=new_member.pseudo, team_name=team.name
    )

    for user in users_to_notify:
        notification.add_all_player_ids(user, kwargs)
        if user.id not in users_ids:
            users_ids.append(user.id)
    send_notification.queue(
        notification_id=notification.id,
        users_ids=users_ids,
        contents=contents,
        **kwargs
    )

def send_notification(
    notification_id,
    contents,
    *,
    headings=None,
    filters=None,
    users_ids=None,
    delivery_time_of_day=None,
    add_history=True,
    **kwargs
):
    if current_app.config["TESTING"]:
        current_app.logger.info("TESTING")

        return

    notification = Notification.query.get(notification_id)

    challenge = Challenge.query.get(kwargs.get("challenge_id"))

    if challenge and notification not in challenge.notifications:
        return

    g.lang = kwargs.get("lang")
    refresh()

    all_attributes_template_id = notification.get_attributes_starting_with(
        "template_id"
    )

    history_created = False
    for attribute in all_attributes_template_id:
        template_id = getattr(notification, attribute)

        if not template_id:
            continue
        kind_app = attribute.split("template_id_")[1]

        player_ids = kwargs.get("player_ids_{}".format(kind_app))

        api_key = current_app.config.get(
            "ONESIGNAL_API_KEY_{}".format(kind_app.upper())
        )
        app_id = current_app.config.get("ONESIGNAL_APP_ID_{}".format(kind_app.upper()))

        header = {
            "Content-Type": "application/json; charset=utf-8",
            "Authorization": "Basic {}".format(api_key),
        }
        payload = {"app_id": app_id, "template_id": template_id, "contents": contents}
        if headings:
            payload["headings"] = headings
        if player_ids:
            payload["include_player_ids"] = player_ids
        elif filters:
            payload["filters"] = filters
            payload["included_segments"] = ["All"]
        else:
            if not history_created and add_history:
                for user_id in users_ids:
                    History.create(
                        notification_id=notification_id,
                        user_id=user_id,
                        bodies=contents,
                        challenge_id=kwargs.get("challenge_id"),
                        team_id=kwargs.get("team_id"),
                    )
                history_created = True
            continue
        if delivery_time_of_day:
            payload["delayed_option"] = "timezone"
            payload["delivery_time_of_day"] = delivery_time_of_day
        t = send_using_onesignal(header, payload)
        description = t.text if t else None
        for user_id in users_ids:
            if not history_created and add_history:
                History.create(
                    notification_id=notification_id,
                    user_id=user_id,
                    bodies=contents,
                    description=description,
                    challenge_id=kwargs.get("challenge_id"),
                    team_id=kwargs.get("team_id"),
                    sent=True,
                )
            history_created = True

def send_using_onesignal(header, payload):
    from app.users.models.user import User

    t = requests.post(
        "https://onesignal.com/api/v1/notifications",
        headers=header,
        data=json.dumps(payload),
    )
    data = t.json()
    errors = data.get("errors")
    if t.status_code != 200 or errors:
        if type(errors) == dict:
            for player_id in errors.get("invalid_player_ids"):
                user = User.query.filter(
                    or_(
                        User.onesignal_ids_apps.any(player_id),
                        User.onesignal_ids_challenge_apps.any(player_id),
                        User.onesignal_ids_web.any(player_id),
                    )
                ).first()
                if user.onesignal_ids_web and player_id in user.onesignal_ids_web:
                    user.onesignal_ids_web.remove(player_id)
                elif (
                    user.onesignal_ids_challenge_apps
                    and player_id in user.onesignal_ids_challenge_apps
                ):
                    user.onesignal_ids_challenge_apps.remove(player_id)
                elif user.onesignal_ids_apps and player_id in user.onesignal_ids_apps:
                    user.onesignal_ids_apps.remove(player_id)
                db.session.commit()
        else:
            current_app.logger.warning(
                """Error from Onesignal : {}\n,
                Payload : {}\n,
                Headers: {}\n
                """.format(
                    errors, data, payload
                )
            )
    else:
        current_app.logger.info(t.text)



```
