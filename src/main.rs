extern crate iron;

use std::io::prelude::*;
use std::fs::File;
use std::path::Path;
use std::error::Error;
use iron::{Iron, Request, Response, IronResult, status};

//use router::Router;

fn main() {
    fn get_file() -> &str {
        let mut path = Path::new("pages/home.html");
        let mut file = match File::open(&path) {
            Ok(file) => file,
            Err(reason) => panic!("couldn't do the thing"),
        };
        let mut s = String::new();
        let mut readFile = match file.read_to_string(&mut s) {
            Ok(file_string) => s,
            Err(reason) => {
                panic!("couldn't do the thing")
            },
        };
    }

    fn hello_world(_: &mut Request) -> IronResult<Response> {
        Response::with((status::Ok, get_file()))
    }

    println!("Listening on port 3000");
    Iron::new(hello_world).http("localhost:3000").unwrap();
}
