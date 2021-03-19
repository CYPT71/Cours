import pymysql

with pymysql.connect(host="127.0.0.1", port=3306, user="root", password="passwd", database="aircraft") as conn:
    cursor = conn.cursor()

    with open('aircraft.sql') as f:
        cursor.execute(f.read().decode('utf-8'), multi=True)

    conn.commit()

    

