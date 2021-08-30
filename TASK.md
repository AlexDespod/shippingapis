Here is the test task details. I would like to see a sample Golang solution that implements these requirements;

Calling multiple API/Providers
The concept is to call several shipping companies' API for their shipping pricing/offers and select the best deal. (Mock API)

Conditions:
No UI expected.
No SQL required.
Few unit tests.

Process Input:
* one set of data {{shipping source address}, {shipping destination address}, [{carton/box dimensions}]}
* Multiple API using the same data with different signatures

Process Output:
* All API will respond with the same data but in different formats
* Process must query, then select the lowest offer and return it in the least amount of time
 
These are sample APIs, each with its own url and credentials and output format
API1 (JSON)
                Input {contact address, warehouse address, package dimensions:[]}
                Output {total}
API2 (JSON)
                Input {consignee, consignor, cartons:[]}
                Output {amount}

API2 (XML)
                Input {consignee, consignor, cartons:[]}
                Output {amount}

