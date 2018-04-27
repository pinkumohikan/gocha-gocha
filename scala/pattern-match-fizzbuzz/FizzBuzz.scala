object FizzBuzz {
  def main(args: Array[String]): Unit = {
    for (i <- 1 to 16) i match {
      case i if i % 3 == 0 && i % 5 == 0 => println("FizzBuzz")
      case i if i % 3 == 0 => println("Fizz")
      case i if i % 5 == 0 => println("Buzz")
      case _ => println(i)
    }
  }
}
