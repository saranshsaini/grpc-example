use arithmetic::arithmetic_service_client::ArithmeticServiceClient;
use arithmetic::AddRequest;

pub mod arithmetic {
    tonic::include_proto!("arithmetic");
}


#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = ArithmeticServiceClient::connect("http://34.123.142.16:50051").await?;

    let request = tonic::Request::new(AddRequest {
        a: 1,
        b: 2,
    });

    let response = client.add(request).await?;

    println!("RESPONSE={:?}", response);

    Ok(())
}   