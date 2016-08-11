
import scala.concurrent.duration._

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import io.gatling.jdbc.Predef._

class Dashboard extends Simulation {

    val httpProtocol = http
        .baseURL("http://localhost:3002")
        .inferHtmlResources()
        .acceptHeader("*/*")
        .acceptEncodingHeader("gzip, deflate, sdch")
        .acceptLanguageHeader("en-US,en;q=0.8,nl;q=0.6")
        .userAgentHeader("Stressy")

    val feeder = Iterator.continually(Map(
        "username" -> "tim",
        "password" -> "ooohthisissosecret",
        "accessToken" -> "F_kd3g5iShGHbB9DvvZOMg"
    ))

    val scn = scenario("Dashboard")
        .feed(feeder)
        /*.exec(http("Get token")
            .post("/token")
            .headers(Map(
                "Authorization" -> "Basic dGltOm9vb2h0aGlzaXNzb3NlY3JldA==",
                "Content-Type" -> "application/x-www-form-urlencoded"
            ))
            .body(StringBody("" +
                 "grant_type=password&username=${username}&password=${password}&scope=profile" +
             ""))
            .check(
                jsonPath("$.access_token").saveAs("accessToken")
            )
        )
        .pause(1)*/
        .exec(http("Popular comics")
            .get("/comic/get_popular")
            .headers(Map(
                "Authorization" -> "Bearer ${accessToken}"
            ))
        )
        .pause(0)
        .exec(http("Awesome stats")
            .get("/awesome_stats")
            .headers(Map(
                "Authorization" -> "Bearer ${accessToken}"
            ))
        )
        .pause(0)
        .exec(http("Notifications")
            .get("/comic/notifications")
            .headers(Map(
                "Authorization" -> "Bearer ${accessToken}"
            ))
        )
        .pause(0)
        .exec(http("Comic publish activity")
            .get("/comic/get_publish_activity")
            .headers(Map(
                "Authorization" -> "Bearer ${accessToken}"
            ))
        )
        .pause(0)
        .exec(http("Trending comics")
            .get("/comic_history/get_trending")
            .headers(Map(
                "Authorization" -> "Bearer ${accessToken}"
            ))
        )

    setUp(scn.inject(rampUsers(3000) over (60 seconds))).protocols(httpProtocol)
    // setUp(scn.inject(atOnceUsers(1))).protocols(httpProtocol)
}
