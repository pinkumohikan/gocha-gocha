package example

import java.time.ZonedDateTime

import software.amazon.awssdk.regions.Region
import software.amazon.awssdk.services.sqs.SqsClient
import software.amazon.awssdk.services.sqs.model.SendMessageRequest

object Hello extends  App {
  val count = 10000
  println("started")
  val start = ZonedDateTime.now()

  val sqsClient = SqsClient.builder().region(Region.AP_NORTHEAST_1).build()
  val queueUrl = sys.env("QUEUE_URL")

  (1 to count)
    .map(id => SendMessageRequest.builder().queueUrl(queueUrl).messageBody(s"id=$id").build())
    .map(req => sqsClient.sendMessage(req))

  println("finished")
  println(s"elapsed: ${ZonedDateTime.now().toEpochSecond - start.toEpochSecond} sec")
}
