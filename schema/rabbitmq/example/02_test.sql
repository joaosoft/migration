
-- migrate up
{
  "queues": [
    {
      "vhost": "dev",
      "name": "new-queue-2",
      "durable": true,
      "auto_delete": false,
      "arguments": {
        "x-dead-letter-exchange": "new-queue-2-exchange-dead",
        "x-dead-letter-routing-key": "new-queue-2-routing-dead"
      }
    }
  ],
  "exchanges": [
    {
      "vhost": "dev",
      "name": "new-queue-2-exchange",
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
      "source": "new-queue-2-exchange",
      "destination": "new-queue-2",
      "destination_type": "queue",
      "routing_key": "example-2",
      "arguments": {}
    }
  ]
}


-- migrate down
