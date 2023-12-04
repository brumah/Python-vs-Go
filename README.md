# Python-vs-Go

database-creation.ipynb
    -db_params need to be altered for each users database connection
    -the chunks of code in this file that begin with "Generate" need to have the num_records... variable altered depending on how many records you want added to that table

servers
    -main.go connects to localhost port 8080
    -main.py connects to localhost port 8000

query requests
    -0 join: http://localhost:8080/getquery?id=1 or http://localhost:8000/getquery?id=1
    -1 join: http://localhost:8080/getquery?id=2 or http://localhost:8000/getquery?id=2
    -2 join: http://localhost:8080/getquery?id=3 or http://localhost:8000/getquery?id=3

wrk commands (port number in command depends on whether the Go or Python server is running)
    -experiment 1 (no concurrency)
        -wrk -t1 -c1 -d30 "http://localhost:8000/getquery?id=1" 
        -wrk -t1 -c1 -d30 "http://localhost:8000/getquery?id=2"
        -wrk -t1 -c1 -d30 "http://localhost:8000/getquery?id=3"
    -experiment 2 (concurrency)
        -wrk -t5 -c5 -d30 "http://localhost:8000/getquery?id=1"
        -wrk -t5 -c5 -d30 "http://localhost:8000/getquery?id=2"
        -wrk -t5 -c5 -d30 "http://localhost:8000/getquery?id=3"