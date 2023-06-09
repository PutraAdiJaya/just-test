# just-test
just test saja

Running Beckend :
 - go mod download
 - go run api.go
 - sample request : 
        - POST : localhost:8080/production-cost
            - payload :
                    {
                        "snaki": {
                            "sugar": 30,
                            "tapioca": 120,
                            "chocolate": 40,
                            "packaging": 1
                        },
                        "cost": {
                            "sugar": 10000,
                            "tapioca": 25000,
                            "chocolate": 35000,
                            "packaging": 500
                        }
                    }
            - response : 
                    {
                        "production_cost": 5200
                    }


