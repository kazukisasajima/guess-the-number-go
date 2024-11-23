package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInt() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n') // 標準入力から1行読み取る
		if err != nil {
			fmt.Println("入力エラー:", err)
			continue
		}

		// 改行や余分な空白を削除
		cleanedInput := strings.TrimSpace(input)

		// 数値に変換
		num, err := strconv.Atoi(cleanedInput)
		if err != nil {
			fmt.Println("無効な入力です。整数を入力してください。")
			continue
		}
		return num
	}
}



func main() {
	var n, m int

	// 入力を受け付ける
	for {
		// 最小値を入力
		fmt.Println("最小値を入力してください: ")
		n = readInt()

		// 最大値を入力
		fmt.Print("最大値を入力してください: ")
		m = readInt()

		// 最小値が最大値以下か確認
		if n < m {
			break
		}
		fmt.Println("最大値は最小値より大きい数値を入力してください！\n")
	}

	// nからmの範囲内で乱数を生成
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	targetNumber := rng.Intn(m-n+1) + n
	maxAttempts := 5

	fmt.Printf("\n%dから%dまでの数字を当ててください！試行回数は%d回です。\n", n, m, maxAttempts)

	for i := 1; i <= maxAttempts; i++ {
		fmt.Printf("%d回目: 予想した数字を入力してください: ", i)
		guess := readInt()

		if guess > targetNumber {
			fmt.Println("Too high!")
		} else if guess < targetNumber {
			fmt.Println("Too low!")
		} else {
			fmt.Println("正解です！おめでとうございます！")
			return
		}
	}

	fmt.Printf("\n試行回数が終了しました！正解は%dでした。\n", targetNumber)
	return
}