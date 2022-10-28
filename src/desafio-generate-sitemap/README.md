https://github.com/diazjf/go-templating/blob/master/example-3/zoo.go

Template index

```xml
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
>
    <sitemap>
        <loc>https://www.jaccede.com/sitemap_pages.xml.gz</loc>
    </sitemap>
    {% for i in range(0,total_sitemap_place) %}
        <sitemap>
            <loc>https://www.jaccede.com/sitemap_places_{{ i + 1 }}.xml.gz
            </loc>
        </sitemap>
    {% endfor %}
    {% for i in range(0,total_sitemap_user) %}
        <sitemap>
            <loc>https://www.jaccede.com/sitemap_users_{{ i + 1 }}.xml.gz</loc>
        </sitemap>
    {% endfor %}
</sitemapindex>
```

Template pages and custom objects

```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
        xmlns:xhtml="http://www.w3.org/1999/xhtml">
    {% if home %}
        {% for lang in languages %}
            <url>
                <loc>https://www.jaccede.com/{{ lang }}/</loc>
                {% for lang in languages %}
                    <xhtml:link
                            rel="alternate"
                            hreflang="{{ lang }}"
                            href="https://www.jaccede.com/{{ lang }}/"
                    />
                {% endfor %}
                <xhtml:link rel="alternate" href="https://www.jaccede.com/fr/"
                            hreflang="x-default"/>
            </url>
        {% endfor %}
        {% for lang in languages %}
            <url>
                <loc>https://www.jaccede.com/{{ lang }}/p/s</loc>
                {% for lang in languages %}
                    <xhtml:link
                            rel="alternate"
                            hreflang="{{ lang }}"
                            href="https://www.jaccede.com/{{ lang }}/p/s"
                    />
                {% endfor %}
                <xhtml:link rel="alternate"
                            href="https://www.jaccede.com/fr/p/s"
                            hreflang="x-default"/>
            </url>
        {% endfor %}
    {% endif %}



    {% for object in objects %}
        {% for lang in languages %}
            <url>
                <loc>https://www.jaccede.com/{{ lang }}/{{ object[0] }}</loc>
                <lastmod>{{ object[1] }}</lastmod>
                {% for lang in languages %}
                    <xhtml:link
                            rel="alternate"
                            hreflang="{{ lang }}"
                            href="https://www.jaccede.com/{{ lang }}/{{ object[0] }}"
                    />
                {% endfor %}
                <xhtml:link rel="alternate"
                            href="https://www.jaccede.com/fr/{{ object[0] }}"
                            hreflang="x-default"/>
            </url>

        {% endfor %}
    {% endfor %}
</urlset>
```

```python
def sitemap():
    from lxml import etree
    from app.places.models.place import Place
    from app.cms.models import Page
    from app.users.models.user import User

    objects = []
    files = []
    last_website_update = datetime.datetime(2017, 1, 20)

    pages = (
        Page.query.filter(Page.slug.notilike("accueil")).order_by(Page.modified).all()
    )

    for page in pages:
        if page.is_main_menu:
            url = "a/{}".format(page.slug)
        else:
            url = "page/{}".format(page.slug)
        if last_website_update > page.modified:
            modified = last_website_update.date().isoformat()
        else:
            modified = page.modified.date().isoformat()
        objects.append([url, modified])
    sitemap_xml = render_template(
        "sitemap.xml",
        objects=objects,
        languages=list(current_app.config["TRANSLATION_LANGS"]),
        home=True,
        modified=datetime.date.today().isoformat(),
    )
    with open("sitemap_temp.xml", "w") as sitemap:
        sitemap.write(sitemap_xml)
        parser = etree.XMLParser(remove_blank_text=True)
        tree = etree.parse(r"sitemap_temp.xml", parser=parser)
        root = tree.getroot()
        etree.ElementTree(root).write("sitemap_pages.xml", method="xml")
    os.remove("sitemap_temp.xml")
    check_call(
        [
            "gzip",
            "-f",
            "{}/sitemap_pages.xml".format(current_app.config["ROOT_PROJECT"]),
        ]
    )
    files.append("sitemap_pages.xml.gz")

    total_places = (
        db.session.query(count(Place.id))
        .filter(
            Place.permanently_closed.isnot(True),
            Place.google_place_id.isnot(None),
            Place.canonical_id.is_(None),
        )
        .scalar()
    )
    batch_size = 4000  # Max of object by xml (4 url by object)
    rounds = int(math.ceil(total_places / batch_size))
    for i in range(0, rounds):
        places = (
            Place.query.filter(
                Place.permanently_closed.isnot(True),
                Place.google_place_id.isnot(None),
                Place.canonical_id.is_(None),
            )
            .order_by(desc(Place.modified))
            .offset(i * batch_size)
            .limit(batch_size)
            .all()
        )
        objects = []
        for place in places:
            url = "p/{}/{}?page=1".format(place.google_place_id, place.slug)
            if last_website_update > place.modified:
                modified = last_website_update.date().isoformat()
            else:
                modified = place.modified.date().isoformat()
            objects.append([url, modified])
        sitemap_xml = render_template(
            "sitemap.xml",
            objects=objects,
            languages=list(current_app.config["TRANSLATION_LANGS"]),
        )
        with open("sitemap_temp.xml", "w") as sitemap:
            sitemap.write(sitemap_xml)
            parser = etree.XMLParser(remove_blank_text=True)
            tree = etree.parse(r"sitemap_temp.xml", parser=parser)
            root = tree.getroot()
            etree.ElementTree(root).write(
                "sitemap_places_{}.xml".format(i + 1), method="xml"
            )
        os.remove("sitemap_temp.xml")
        check_call(
            [
                "gzip",
                "-f",
                "{}/sitemap_places_{}.xml".format(
                    current_app.config["ROOT_PROJECT"], i + 1
                ),
            ]
        )
        files.append("sitemap_places_{}.xml.gz".format(i + 1))

    total_users = (
        db.session.query(count(User.id))
        .filter(
            User.is_activated.is_(True),
            User.is_blocked.is_(False),
            User.is_suspended.is_(False),
        )
        .scalar()
    )
    batch_size = 4000  # Max of object by xml (4 url by object)
    rounds_user = int(math.ceil(total_users / batch_size))
    for i in range(0, rounds_user):
        users = (
            User.query.filter(
                User.is_activated.is_(True),
                User.is_blocked.is_(False),
                User.is_suspended.is_(False),
            )
            .order_by(desc(User.modified))
            .offset(i * batch_size)
            .limit(batch_size)
            .all()
        )
        objects = []
        for user in users:
            modified = user.modified if user.modified else user.created
            if last_website_update > modified:
                modified = last_website_update.date().isoformat()
            else:
                modified = modified.date().isoformat()
            url_contribution = "u/{}/contributions?page=1".format(user.slug)
            objects.append([url_contribution, modified])

        sitemap_xml = render_template(
            "sitemap.xml",
            objects=objects,
            languages=list(current_app.config["TRANSLATION_LANGS"]),
        )
        with open("sitemap_temp.xml", "w") as sitemap:
            sitemap.write(sitemap_xml)
            parser = etree.XMLParser(remove_blank_text=True)
            tree = etree.parse(r"sitemap_temp.xml", parser=parser)
            root = tree.getroot()
            etree.ElementTree(root).write(
                "sitemap_users_{}.xml".format(i + 1), method="xml"
            )
        os.remove("sitemap_temp.xml")
        check_call(
            [
                "gzip",
                "-f",
                "{}/sitemap_users_{}.xml".format(
                    current_app.config["ROOT_PROJECT"], i + 1
                ),
            ]
        )
        files.append("sitemap_users_{}.xml.gz".format(i + 1))

    sitemap_xml = render_template(
        "sitemap_index.xml", total_sitemap_place=rounds, total_sitemap_user=rounds_user
    )
    with open("sitemap.xml", "w") as sitemap:
        sitemap.write(sitemap_xml)
    # parser = etree.XMLParser(remove_blank_text=True)
    #     tree = etree.parse(r'sitemap_temp.xml',parser=parser)
    #     root = tree.getroot()
    #     etree.ElementTree(root).write('sitemap.xml',method='xml')
    files.append("sitemap.xml")

    check_call(["tar", "-cf", "sitemap.tar"] + files)
    check_call(["rm"] + files)

```
