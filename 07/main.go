package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arvindkhoisnam/challenges/07/agent"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func startMetricsServer() {
    http.Handle("/metrics", promhttp.Handler())
    go func() {
        if err := http.ListenAndServe(":2112", nil); err != nil {
            log.Printf("Error starting metrics server: %v", err)
        }
    }()
}

func main(){
	startMetricsServer()
	
	// go func ()  {
	// 	client := agent.NewAgent()
	// 	client.ExecuteTask(`Go to http://localhost:5173, fill out the contact form with name: Misha Mansoor, email: misha@example.com, contact: 9876543210,
	// 	hobbies: Playing djenty riffs and message: You're a slave born into a dark world of deceit.
	// 	image: https://promixacademy.com/wp-content/uploads/2021/10/Misha-Mansoor_01-1024x632.jpeg
	// 	and submit it`)

	// defer client.Close()
	// fmt.Scanln()
	// }()
	// go func ()  {
	// 	client := agent.NewAgent()
	// 	client.ExecuteTask(`Go to http://localhost:5173, fill out the contact form with name: Dave Mustaine, email: davemustaine@example.com, contact: 9876543210,
	//     hobbies: Playing heavy riffs and message: Hello me, meet the real me!
	//     image: https://gazette.gibson.com/wp-content/uploads/2024/06/Dave-Mustaine-strings-and-strap-2024_gazette_v3.jpg
	// 	and submit it`)

	// defer client.Close()
	// fmt.Scanln()
	// }()
	// go func ()  {
	// 	client := agent.NewAgent()
	// 	client.ExecuteTask(`Go to http://localhost:5173, fill out the contact form with name: Adam Jones, email: adamjones@example.com, contact: 9876543210,
	// 	hobbies: Playing chunky riffs and message: Constant over stimulation numbs me, But I would not want you any other way.
	// 	image: https://images2.minutemediacdn.com/image/upload/c_crop,x_0,y_253,w_1997,h_1123/c_fill,w_1080,ar_16:9,f_auto,q_auto,g_auto/images%2FGettyImages%2Fmmsport%2F345%2F01hy448fpxktbfb776re.jpg
	// 	and submit it`)
	// defer client.Close()
	// fmt.Scanln()
	// }()
	client := agent.NewAgent()
	client.ExecuteTask(`Go to http://localhost:5173, fill out the contact form with name: Adam Jones, email: adamjones@example.com, contact: 9876543210,
	hobbies: Playing chunky riffs and message: Constant over stimulation numbs me, But I would not want you any other way.
	image: https://images2.minutemediacdn.com/image/upload/c_crop,x_0,y_253,w_1997,h_1123/c_fill,w_1080,ar_16:9,f_auto,q_auto,g_auto/images%2FGettyImages%2Fmmsport%2F345%2F01hy448fpxktbfb776re.jpg
	and submit it`)
	defer client.Close()
	fmt.Scanln()
}

