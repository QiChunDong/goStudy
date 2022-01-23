package pract

import (
  t "gostudy/tools"
  "math"
)

func ClacBMI(W float64, H float64)  {
  // math.Pow 取方 第二个参数指定方次
  BMI := W / math.Pow(H, 2)
  t.Printfln("BMI: %.2f", BMI)

  if BMI < 18.5 {
    t.Printfln("偏廋")
  } else if (BMI < 24) && (18.5 <= BMI) {
    t.Printfln("正常")
  } else if (BMI < 28) && (24 <= BMI) {
    t.Printfln("偏胖")
  } else if (BMI < 30) && (28 <= BMI) {
    t.Printfln("肥胖")
  } else {
    t.Printfln("重度肥胖")
  }
}