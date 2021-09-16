# API Contract

**running port = 127.0.0.1:8080**

* Get Order <br>
    * http://127.0.0.1:8080/orders/(id) method : get
* Get Orders <br>
    * http://127.0.0.1:8080/orders/  method : get
* Create Order
    * http://127.0.0.1:8080/orders method : post
    * Request body
----
  ```json
{
  "orderedAt": "2019-11-09T21:21:46+00:00",
  "customerName": "Tome Halldan",
  "items": [
    {
      "itemCode": "888",
      "description": "Iphone 13",
      "quantity": 1
    }
  ]
}
```
* Update Order
  * http://127.0.0.1:8080/orders method : put
  * Request body
----
  ```json
{
   "orderId":3,
   "customerName":"Thomas Marc",
   "orderedAt":"2019-11-09T21:21:46Z",
   "items":[
      {
         "lineItemId":1,
         "itemCode":"333",
         "description":"IPhone 7",
         "quantity": 4
      }
   ]
}
```
* Delete order
  * http://127.0.0.1:8080/orders/(id) method delete