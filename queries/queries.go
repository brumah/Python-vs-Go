package queries

import (
	"database/sql"
	"fmt"
	"log"
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

	query := "SELECT username, date_created, follower_count, post_count FROM user_data;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	swapHTML := `<table class='table-auto w-full border border-collapse text-center'>
	<thead>
	<tr>
	<th class='px-4 py-2 border border-gray-600'>Username</th>
	<th class='px-4 py-2 border border-gray-600'>Account Creation Date</th>
	<th class='px-4 py-2 border border-gray-600'>Follower Count</th>
	<th class='px-4 py-2 border border-gray-600'>Post Count</th>
	</tr>
	</thead>
	<tbody>`

	count := 0
	for rows.Next() && count <= 4 {
		var column1 string
		var column2 time.Time
		var column3, column4 int
		if err := rows.Scan(&column1, &column2, &column3, &column4); err != nil {
			log.Fatal(err)
		}
		swapHTML = swapHTML + fmt.Sprintf(`
		<tr>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%v</td>
		<td class='px-4 py-2 border border-gray-600'>%v</td>
		</tr>`,
			column1, column2.Format("2006-01-02"), column3, column4)
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

	query := "SELECT username, date_created, follower_count, post_count FROM user_data WHERE follower_count < 10000;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	swapHTML := `<table class='table-auto w-full border	border-gray-600 text-center'>
	<thead>
	<tr>
	<th class='px-4 py-2 border-separate border-gray-600'>Username</th>
	<th class='px-4 py-2 border-separate border-gray-600'>Account Creation Date</th>
	<th class='px-4 py-2 border-separate border-gray-600'>Follower Count</th>
	<th class='px-4 py-2 border-separate border-gray-600'>Post Count</th>
	</tr>
	</thead>
	<tbody>`

	count := 0
	for rows.Next() && count <= 4 {
		var column1 string
		var column2 time.Time
		var column3, column4 int
		if err := rows.Scan(&column1, &column2, &column3, &column4); err != nil {
			log.Fatal(err)
		}
		swapHTML = swapHTML + fmt.Sprintf(`
		<tr>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%v</td>
		<td class='px-4 py-2 border border-gray-600'>%v</td>
		</tr>`,
			column1, column2.Format("2006-01-02"), column3, column4)
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
				<th class="px-4 py-2 border border-gray-600">Follower Count</th>
				<th class="px-4 py-2 border border-gray-600">Post Count</th>
			</tr>
		</thead>
	</table>
	`
}
