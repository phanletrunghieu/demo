for x in ./migrator/*.json; do
    aws dynamodb --endpoint-url http://localhost:8000 create-table --cli-input-json file://$x
done