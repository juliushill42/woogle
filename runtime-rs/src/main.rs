use tokio::time::{sleep, Duration};

#[tokio::main]
async fn main() {
    loop {
        println!("WOOGLE runtime heartbeat");
        sleep(Duration::from_secs(5)).await;
    }
}