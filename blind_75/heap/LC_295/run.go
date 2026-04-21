package main

import "github.com/sirupsen/logrus"

func main() {
	mf := Constructor()

	mf.AddNum(1)
	logrus.Printf("%f", mf.FindMedian()) // 1.0

	mf.AddNum(3)
	logrus.Printf("%f", mf.FindMedian()) // 2.0

	mf.AddNum(2)
	logrus.Printf("%f", mf.FindMedian()) // 2.0
}