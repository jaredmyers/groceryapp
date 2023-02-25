import mysql.connector


def main():

    conn = mysql.connector.connect(
                host="172.19.255.201",
                user="root",
                password="12345",
                database="user_service"
            )

    print("connected to mysql")

    process_login(conn)


def process_login(conn):
    '''check user provided login creds against db'''

    username = "demo"
    # password = "password"

    query = "select uname, pw from users where uname=%s;"
    val = (username,)
    cursor = conn.cursor()
    cursor.execute(query, val)
    query_result = cursor.fetchall()

    if not query_result:
        return

    uname = query_result[0][0]
    pw = query_result[0][1]

    print("results from database:")
    print(uname, pw)


main()
