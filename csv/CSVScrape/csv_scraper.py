from csv import DictReader
from datetime import datetime

with open("data/nfl-2019-pittsburgh-steelers-GMTStandardTime.csv") as document:
    home = DictReader(document)
    print("Home Game with Away Team on Date")
    for row in home:
        day = datetime.strptime(row.get("Date"), "%d/%m/%Y %H:%M")
        if "Heinz Field" in row.get("Location"):
            print("Home Game with %s on %s" % (row.get("Away Team"), row.get("Date")))

with open("data/nfl-2019-pittsburgh-steelers-GMTStandardTime.csv") as document:
    away = DictReader(document)
    print()
    print("Away Game to Home Team on Date")
    for row in away:
        day = datetime.strptime(row.get("Date"), "%d/%m/%Y %H:%M")
        if "Heinz Field" not in row.get("Location"):
            print(
                "Away Game to %s at %s on %s"
                % (row.get("Home Team"), row.get("Location"), row.get("Date"))
            )
