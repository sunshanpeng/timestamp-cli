package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/sunshanpeng/timestamp-cli/pkg/converter"
	"github.com/sunshanpeng/timestamp-cli/pkg/formatter"
	"github.com/sunshanpeng/timestamp-cli/pkg/parser"
)

var (
	useUTC      bool
	useSecond   bool
	useMS       bool
	showTZ      bool
	showVersion bool
)

const version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "timestamp [input]",
	Short: "快速处理时间戳的命令行工具",
	Long: `timestamp 是一个用于快速处理时间戳的命令行工具。

支持功能：
  - 获取当前时间戳（秒级/毫秒级）
  - 时间戳转日期字符串
  - 日期字符串转时间戳
  - UTC 时区支持

示例：
  timestamp                           # 显示完整信息
  timestamp -s                        # 仅输出秒级时间戳
  timestamp -ms                       # 仅输出毫秒级时间戳
  timestamp 1768809600                # 时间戳转日期
  timestamp "2026-01-19 16:00:00"     # 日期转时间戳
  timestamp -utc                      # UTC 时区完整信息
  timestamp -tz                       # 显示时区信息`,
	Run: run,
}

func init() {
	rootCmd.Flags().BoolVar(&useUTC, "utc", false, "使用 UTC 时区")
	rootCmd.Flags().BoolVarP(&useSecond, "second", "s", false, "仅输出秒级时间戳")
	rootCmd.Flags().BoolVar(&useMS, "ms", false, "输出/使用毫秒级时间戳")
	rootCmd.Flags().BoolVar(&showTZ, "tz", false, "显示时区信息")
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "显示版本信息")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	// 显示版本信息
	if showVersion {
		fmt.Printf("timestamp version %s\n", version)
		return
	}

	// 确定时区
	loc := time.Local
	if useUTC {
		loc = time.UTC
	}

	// 显示时区信息
	if showTZ {
		fmt.Println(formatter.FormatTimezone(loc))
		return
	}

	// 无参数：显示当前时间信息
	if len(args) == 0 {
		// 仅输出秒级时间戳
		if useSecond {
			fmt.Println(converter.GetCurrentTimestamp(loc, false))
			return
		}
		// 仅输出毫秒级时间戳
		if useMS {
			fmt.Println(converter.GetCurrentTimestamp(loc, true))
			return
		}
		// 显示完整信息
		fmt.Println(formatter.FormatFull(loc))
		return
	}

	// 有参数：解析并转换
	input := args[0]
	inputType, t, err := parser.ParseInput(input, loc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n\n", err)
		fmt.Fprintln(os.Stderr, "使用示例:")
		fmt.Fprintln(os.Stderr, "  timestamp 1768809600                # 时间戳转日期")
		fmt.Fprintln(os.Stderr, "  timestamp 1768809600000             # 毫秒级时间戳转日期")
		fmt.Fprintln(os.Stderr, "  timestamp \"2026-01-19 16:00:00\"     # 日期转时间戳")
		os.Exit(1)
	}

	result := converter.Convert(inputType, t, useMS)
	fmt.Println(result)
}
