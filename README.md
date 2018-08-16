# fun-with-chan

The goal of this project is to have fun with ring.Ring, chan, grpc and go routines


## Completed

### Basic operations
- Subscribe
- Put
- Push
- Mark done

## To do

### Timeouts
- Part of newTopic()
- Part of canWork()
- Part of markDone()

### Multiple consumers
- Part of Topic
- Part of canWork()
- Part of work()

### GRPC
- Messages for newTopic(), put(), subscribe(), push(), disconnect
- Client for put()
- Client for subscribe()

### Prettiness
- Web server
- Health check rest endpoint
- Fancy healthy rest endpoint (puts, pushes, count at time, up-time, processing time, ingestion time)
- HTML page for querying the endpoints