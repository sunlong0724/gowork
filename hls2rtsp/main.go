package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

type Hls2RtspItem struct {
	Hls     string `mapstructure:"hls"`
	Rtsp    string `mapstructure:"rtsp"`
	Cmd 	string `mapstructure:"cmd"`
}
type Config struct {
	Items map[string]Hls2RtspItem `mapstructure:"items"`
}

func (this *Hls2RtspItem) String() string {
	return fmt.Sprintf("cmd:%s,hls:%s,rtsp:%s", this.Cmd,this.Hls, this.Rtsp)
}

func main() {
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	viper.SetConfigName("hls2rtsp")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Read hls2rtsp.yaml failed. err :", err)
		panic(err)
	}
	c := &Config{}
	if err := viper.Unmarshal(c); err != nil {
		fmt.Println("unmarshal config failed, err :", err)
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	for k, v := range c.Items {
		go func(ctx context.Context, key string, value Hls2RtspItem) {
			args := value.Hls
			cmd := exec.Command(value.Cmd, args)
			for {
				if err := cmd.Start(); err != nil {
					fmt.Println("start ", key, " failed.")
				}
				if err := cmd.Wait(); err != nil {
					fmt.Println(cmd.String(), " existed,err:", err)
				}
				select {
				case <-ctx.Done():
					return
				default:
					time.Sleep(time.Second)
				}
			}
		}(ctx, k, v)
	}

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT)
	<-sigCh
	cancel()
	time.Sleep(time.Second * 5)
	fmt.Println("main exited. ", time.Now().Format("2006-01-02 15:04:05"))
}
