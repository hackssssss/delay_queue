package common

const (
	PayloadKeyLength = 16
	KeySep           = `_`
)

//redis
const (
	QueueName       = `delay_queue_queue`
	NotifyQueueName = `delay_queue_notify_queue`
	ZsetName        = `delay_queue_zset`

	//zrange(0, DetectStop), http://doc.redisfans.com/sorted_set/zrange.html
	DetectStop = 10

	PublisherPopQueueTimeout = 2 //seconds
)
