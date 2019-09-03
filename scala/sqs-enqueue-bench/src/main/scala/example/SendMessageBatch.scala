package example

import java.time.ZonedDateTime

import software.amazon.awssdk.regions.Region
import software.amazon.awssdk.services.sqs.SqsClient
import software.amazon.awssdk.services.sqs.model.{SendMessageBatchRequest, SendMessageBatchRequestEntry}

object SendMessageBatch extends App {
  println("started")
  val start = ZonedDateTime.now()

  val sqsClient = SqsClient.builder().region(Region.AP_NORTHEAST_1).build()
  val queueUrl = sys.env("QUEUE_URL")
  val count = sys.env("COUNT").toInt

  (1 to count)
    .map(id => SendMessageBatchRequestEntry.builder().id(s"$id").messageBody(s"id=$id").build())
    .grouped(10)
    .toList
    .map(messages => SendMessageBatchRequest.builder().queueUrl(queueUrl).entries(messages: _*).build())
    .par
    .map(req => {
      println("calling sendMessageBatch")
      sqsClient.sendMessageBatch(req)
    })

  println("finished")
  println(s"elapsed: ${ZonedDateTime.now().toEpochSecond - start.toEpochSecond} sec")
}
