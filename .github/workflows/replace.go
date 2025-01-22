connectionString, found := os.LookupEnv("CONFIGURATION__AZURECOSMOSDB__CONNECTIONSTRING")
if !found {
	panic("Azure Cosmos DB for Table account connection string not set.")
}

log.Println("CONNECTION STRING: ", connectionString)

service, err := aztables.NewServiceClient(connectionString, nil)
if err != nil {
	panic(err)
}