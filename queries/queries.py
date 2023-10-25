import psycopg2
import csv

def blue_button():
    db_connection = psycopg2.connect(
        host='localhost',
        dbname="matthewbruner", 
        user="matthewbruner",
        port="5432"
    )
    cursor = db_connection.cursor()

    query = "SELECT username, date_created FROM users;"

    cursor.execute(query)
    rows = cursor.fetchall()

    swapHTML = """
        <table class='table-auto w-full border border-collapse text-center'>
            <thead>
                <tr>
                    <th class='px-4 py-2 border border-gray-600 bg-slate-400'>Username</th>
                    <th class='px-4 py-2 border border-gray-600 bg-slate-400'>Account Creation Date</th>
                </tr>
            </thead>
            <tbody>
    """

    for i in range(5):
        swapHTML = swapHTML + f"""
            <tr>
                <td class='px-4 py-2 border border-gray-600'>{rows[i][0]}</td>
                <td class='px-4 py-2 border border-gray-600'>{rows[i][1]}</td>
            </tr>
        """

    swapHTML = swapHTML + """
            </tbody>
        </table>
    """

    return swapHTML

def green_button():
    db_connection = psycopg2.connect(
        host='localhost',
        dbname="matthewbruner", 
        user="matthewbruner",
        port="5432"
    )
    cursor = db_connection.cursor()

    query = "SELECT user_id, post_date, id FROM posts WHERE post_Date > '2021-01-01';"

    cursor.execute(query)
    rows = cursor.fetchall()

    swapHTML = """
        <table class='table-auto w-full border border-collapse text-center'>
            <thead>
                <tr>
                    <th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>User ID</th>
                    <th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Post Date</th>
                    <th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Post ID</th>
                </tr>
            </thead>
            <tbody>
    """

    for i in range(5):
        swapHTML = swapHTML + f"""
            <tr>
                <td class='px-4 py-2 border border-gray-600'>{rows[i][0]}</td>
                <td class='px-4 py-2 border border-gray-600'>{rows[i][1]}</td>
                <td class='px-4 py-2 border border-gray-600'>{rows[i][2]}</td>
            </tr>
        """

    swapHTML = swapHTML + """
            </tbody>
        </table>
    """

    return swapHTML

def red_button():
    db_connection = psycopg2.connect(
        host='localhost',
        dbname="matthewbruner", 
        user="matthewbruner",
        port="5432"
    )
    cursor = db_connection.cursor()

    query = """
        SELECT u.username, count(u.username) AS follower_count
            FROM users AS u
            INNER JOIN followers as f
            ON u.id = f.user_id
            GROUP BY u.username
            ORDER BY follower_count DESC;
    """

    cursor.execute(query)
    rows = cursor.fetchall()

    swapHTML = """
        <table class='table-auto w-full border border-collapse text-center'>
            <thead>
                <tr>
                    <th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Username</th>
                    <th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Follower Count</th>
                </tr>
            </thead>
            <tbody>
    """

    for i in range(5):
        swapHTML = swapHTML + f"""
            <tr>
                <td class='px-4 py-2 border border-gray-600'>{rows[i][0]}</td>
                <td class='px-4 py-2 border border-gray-600'>{rows[i][1]}</td>
            </tr>
        """

    swapHTML = swapHTML + """
            </tbody>
        </table>
    """

    return swapHTML

def purple_button():
    db_connection = psycopg2.connect(
        host='localhost',
        dbname="matthewbruner", 
        user="matthewbruner",
        port="5432"
    )
    cursor = db_connection.cursor()

    query = """
		SELECT u.username, p.id, count(l.post_id) AS total_likes
			FROM likes AS l
			LEFT JOIN posts AS p
			ON l.post_id = p.id
			LEFT JOIN users AS u
			ON u.id = p.user_id
			GROUP BY p.id, u.username
			ORDER BY total_likes DESC;
    """

    cursor.execute(query)
    rows = cursor.fetchall()

    swapHTML = """
        <table class='table-auto w-full border border-collapse text-center'>
            <thead>
                <tr>
                    <th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Username</th>
                    <th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Post ID</th>
                    <th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Total Likes</th>
                </tr>
            </thead>
            <tbody>
    """

    for i in range(5):
        swapHTML = swapHTML + f"""
            <tr>
                <td class='px-4 py-2 border border-gray-600'>{rows[i][0]}</td>
                <td class='px-4 py-2 border border-gray-600'>{rows[i][1]}</td>
                <td class='px-4 py-2 border border-gray-600'>{rows[i][2]}</td>
            </tr>
        """

    swapHTML = swapHTML + """
            </tbody>
        </table>
    """
    
    return swapHTML

def clear_button():
    return """
        <table class="table-auto w-full border border-collapse text-center" id="target">
            <thead>
                <tr>
                    <th class="px-4 py-2 border border-gray-600">Username</th>
                    <th class="px-4 py-2 border border-gray-600">Account Creation Date</th>
                </tr>
            </thead>
	    </table>
    """

def export_button(latencies):

    path = "latencies/output.csv"
    with open(path, mode='w', newline='') as file:
        writer = csv.writer(file, delimiter=',', quotechar='"', quoting=csv.QUOTE_MINIMAL)
        for latency in latencies:
            writer.writerow([latency])

    return "Export Latencies"