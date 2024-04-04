## Simple REST API with DynamoDB

### Sample cURL Commands

Here I have used a book object with attributes ISBN, Name, Author, and Published Version for this implementation

#### Get Item by ID

To retrieve an item by its ID, you can use the following `curl` command:

```bash
curl --location 'https://i4ym7uvum6.us-east-1.awsapprunner.com/api/book/{isbn}'
Sample:
curl --location 'https://i4ym7uvum6.us-east-1.awsapprunner.com/api/book/9780547928203'
```
This command sends a GET request to the specified URL, retrieving the item with the provided ISBN number (9780547928203).
Replace ISBN to retrieve any inserted book

#### Post an Item

To add a new item, you can use the following curl command:
```bash
curl --location 'https://i4ym7uvum6.us-east-1.awsapprunner.com/api/book' \
--header 'Content-Type: application/json' \
--data '{
    "isbn_number" : "9780399501487",
    "name":"The Two Towers",
    "author" : "Tolkien, J.R.R.",
    "published_version" : 2
}'
```
This command sends a POST request to the specified URL with JSON data containing the details of the book to be added, including ISBN number, name, author, and published version.

Ensure to replace the values in the JSON payload (isbn_number, name, author, published_version) with the details of the book you want to add.

NOTE : isbn_number should be unique for each book record. If you use same isbn in multiple requests, it will update the relevant record for isbn in DB 