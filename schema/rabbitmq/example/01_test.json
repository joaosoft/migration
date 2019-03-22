
-- migrate up
{
  "queues": [
    {
      "vhost": "dev",
      "name": "new-queue-1",
      "durable": true,
      "auto_delete": false,
      "arguments": {
        "x-dead-letter-exchange": "new-queue-1-exchange-dead",
        "x-dead-letter-routing-key": "new-queue-1-routing-dead"
      }
    }
  ],
  "exchanges": [
    {
      "vhost": "dev",
      "name": "new-queue-1-exchange",
      "type": "direct",
      "durable": true,
      "auto_delete": false,
      "internal": false,
      "arguments": {}
    }
  ],
  "bindings": [
    {
      "vhost": "dev",
      "source": "new-queue-1-exchange",
      "destination": "new-queue-1",
      "destination_type": "queue",
      "routing_key": "example-1",
      "arguments": {}
    }
  ]
}


-- migrate down
