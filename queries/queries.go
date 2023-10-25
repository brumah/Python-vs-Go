package queries

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func BlueButton() string {
	connectionString := "user=matthewbruner host=localhost port=5432 dbname=matthewbruner sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	query := "SELECT username, date_created FROM users;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	swapHTML := `<table class='table-auto w-full border border-collapse text-center'>
	<thead>
	<tr>
	<th class='px-4 py-2 border border-gray-600 bg-slate-400'>Username</th>
	<th class='px-4 py-2 border border-gray-600 bg-slate-400'>Account Creation Date</th>
	</tr>
	</thead>
	<tbody>`

	count := 0
	for rows.Next() && count <= 4 {
		var column1 string
		var column2 time.Time
		if err := rows.Scan(&column1, &column2); err != nil {
			log.Fatal(err)
		}
		swapHTML = swapHTML + fmt.Sprintf(`
		<tr>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		</tr>`,
			column1, column2.Format("2006-01-02"))
		count++
	}
	swapHTML = swapHTML + "</tbody></table>"
	return swapHTML
}

func GreenButton() string {
	connectionString := "user=matthewbruner host=localhost port=5432 dbname=matthewbruner sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	query := "SELECT user_id, post_date, id FROM posts WHERE post_Date > '2021-01-01';"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	swapHTML := `<table class='table-auto w-full border	border-gray-600 text-center'>
	<thead>
	<tr>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>User ID</th>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Post Date</th>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Post ID</th>
	</tr>
	</thead>
	<tbody>`

	count := 0
	for rows.Next() && count <= 4 {
		var column1, column3 string
		var column2 time.Time
		if err := rows.Scan(&column1, &column2, &column3); err != nil {
			log.Fatal(err)
		}
		swapHTML = swapHTML + fmt.Sprintf(`
		<tr>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%v</td>
		</tr>`,
			column1, column2.Format("2006-01-02"), column3)
		count++
	}
	swapHTML = swapHTML + "</tbody></table>"
	return swapHTML
}

func RedButton() string {
	connectionString := "user=matthewbruner host=localhost port=5432 dbname=matthewbruner sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	query := `
		SELECT u.username, count(u.username) AS follower_count
			FROM users AS u
			INNER JOIN followers as f
			ON u.id = f.user_id
			GROUP BY u.username
			ORDER BY follower_count DESC;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	swapHTML := `<table class='table-auto w-full border	border-gray-600 text-center'>
	<thead>
	<tr>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Username</th>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Follower Count</th>
	</tr>
	</thead>
	<tbody>`

	count := 0
	for rows.Next() && count <= 4 {
		var column1 string
		var column2 int
		if err := rows.Scan(&column1, &column2); err != nil {
			log.Fatal(err)
		}
		swapHTML = swapHTML + fmt.Sprintf(`
		<tr>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%v</td>
		</tr>`,
			column1, column2)
		count++
	}
	swapHTML = swapHTML + "</tbody></table>"
	return swapHTML
}

func PurpleButton() string {
	connectionString := "user=matthewbruner host=localhost port=5432 dbname=matthewbruner sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	query := `
		SELECT u.username, p.id, count(l.post_id) AS total_likes
			FROM likes AS l
			LEFT JOIN posts AS p
			ON l.post_id = p.id
			LEFT JOIN users AS u
			ON u.id = p.user_id
			GROUP BY p.id, u.username
			ORDER BY total_likes DESC;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	swapHTML := `<table class='table-auto w-full border	border-gray-600 text-center'>
	<thead>
	<tr>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Username</th>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Post ID</th>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Total Likes</th>
	</tr>
	</thead>
	<tbody>`

	count := 0
	for rows.Next() && count <= 4 {
		var column1, column2 string
		var column3 int
		if err := rows.Scan(&column1, &column2, &column3); err != nil {
			log.Fatal(err)
		}
		swapHTML = swapHTML + fmt.Sprintf(`
		<tr>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%v</td>
		</tr>`,
			column1, column2, column3)
		count++
	}
	swapHTML = swapHTML + "</tbody></table>"
	return swapHTML
}

func ClearButton() string {
	return `
	<table class="table-auto w-full border border-collapse text-center" id="target">
		<thead>
			<tr>
				<th class="px-4 py-2 border border-gray-600">Username</th>
				<th class="px-4 py-2 border border-gray-600">Account Creation Date</th>
			</tr>
		</thead>
	</table>
	`
}

func ExportButton(latencies []string) string {
	latencyValues := latencies

	// Create a new CSV file
	file, err := os.Create("latencies/output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the data slice to the CSV file
	err = writer.Write(latencyValues)
	if err != nil {
		panic(err)
	}
	return "Export Latencies"
}
