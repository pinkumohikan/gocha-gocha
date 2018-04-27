object FizzBuzz {
  def main(args: Array[String]): Unit = {
    for (i <- 1 to 16) (i % 3, i % 5) match {
      case (0, 0) => println("FizzBuzz")
      case (0, _) => println("Fizz")
      case (_, 0) => println("Buzz")
      case _ => println(i)
    }
  }
}
