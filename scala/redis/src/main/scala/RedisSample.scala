import com.redis.RedisClient
import scala.util.control.Exception.catching

object RedisSample {
  def main(args: Array[String]): Unit = {
    val redisClient = new RedisClient("localhost", 6379)

    redisClient.set("hello", "world")
    redisClient.get("hello") match {
      case Some(v) =>
        println(s"hello $v")
      case None =>
        println("あれれー、値がsetできてないよー？")
    }
  }
}
