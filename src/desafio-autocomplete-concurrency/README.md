#

Criar uma busca de places de acordo com uma string pesquisada pelo usuário.

Primeiro pesquisar na base do redis se existe algo relacionado com aquele place

Senão utilizar a api do gmaps places passando

Os resultados encontrados salvar em uma base redis.

---

```
class AutocompleteSchema(ma.Schema):
    jaccede = fields.Boolean(missing=True)
    google = fields.Boolean(missing=True)
    token = fields.String(required=True)
    latitude = fields.Decimal()
    longitude = fields.Decimal()
    radius = fields.Integer(missing=10000)
    types = fields.String()
    components = fields.List(fields.String())
    strict_bounds = fields.Boolean(missing=False)
    session_token = fields.String(required=True)


class AutocompleteThread(Thread):
    def __init__(self, token, enabled):
        Thread.__init__(self)
        self.token = token
        self.enabled = enabled
        self.redis_host = current_app.config["REDISEARCH_HOST"]
        self.redis_port = current_app.config["REDISEARCH_PORT"]
        self.results = []

    def run(self):
        if self.enabled:
            try:
                rs_client = Client(
                    "autocomplete", host=self.redis_host, port=self.redis_port
                )
                for place in rs_client.search(self.token).docs:
                    place_data = loads(place.document)
                    place_data["is_place"] = True
                    self.results.append(place_data)
            except Exception as e:
                current_app.logger.warning(
                    "Error on internal autocomplete search : {}".format(e)
                )


AutocompleteThread(Thread):
    def __init__(self, token, location, data, types, lang, enabled):
        Thread.__init__(self)
        self.token = token
        self.location = location
        self.data = data
        self.types = types
        self.lang = lang
        self.enabled = enabled
        self.results = []

    def run(self):
        if self.enabled:
            try:
                self.results = gmaps.places_autocomplete(
                    input_text=self.token,
                    location=self.location,
                    radius=self.data.get("radius"),
                    types=self.types,
                    components=self.data.get("components"),
                    strict_bounds=self.data.get("strict_bounds"),
                    session_token=self.data.get("session_token"),
                    language=self.lang,
                )
            except Exception as e:
                current_app.logger.warn(
                    "Error on google autocomplete search : {}".format(e)
                )

def index(self):
        data = AutocompleteSchema().load(request.args)
        location = None
        token = unidecode(data.get("token")).lower()
        types = data.get("types")

        if data.get("latitude") and data.get("longitude"):
            location = [data.get("latitude"), data.get("longitude")]

        jaccede_search = data.get("jaccede") and not types or types == "establishment"
        jaccede_thread = JaccedeAutocompleteThread(token, jaccede_search)
        jaccede_thread.start()

        google_search = data.get("google")
        google_thread = GoogleAutocompleteThread(
            token, location, data, types, g.lang, google_search
        )
        google_thread.start()

        jaccede_thread.join()
        google_thread.join()

        results = []
        # For troubleshooting purpose, should be removed once clean solution
        # defined (search for replacement for secondary_text)
        for result in google_thread.results[:5]:
            cats = result.get("types", [])

            if (
                any(elem in cats for elem in NO_MORE_AVAILABLE_CATEGORIES)
                and len(cats) <= 3
            ):
                continue
            results.append(
                {
                    "gpid": result["place_id"],
                    "name": result["structured_formatting"]["main_text"],
                    "full_address": result["structured_formatting"].get(
                        "secondary_text", ""
                    ),
                    "is_place": "establishment" in cats,
                }
            )
        place_set = {x["gpid"] for x in results}
        for result in jaccede_thread.results:
            if result["gpid"] not in place_set:
                results.append(result)
            if len(results) >= 10:
                break

        return PlaceSuggestionSchema(many=True).jsonify(results)

```
