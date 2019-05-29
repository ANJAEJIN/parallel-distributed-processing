use std::io;
use std::error::Error;
use std::fs::File;
use std::path::Path;
use std::io::{Read, BufReader, BufRead};
use std::time::SystemTime;

fn main() {
    let _path = Path::new("./random.txt");
    let _display = _path.display();

    let mut file = File::open(&_path).unwrap();
    let reader = BufReader::new(file);

    let mut data = Vec::new();
    for (index, line) in reader.lines().enumerate() {
        let line = line.unwrap();
        data.push(line.parse::<usize>().unwrap());
    }

    let now = SystemTime::now();
    let mut stack = merge_sort(data);
    match now.elapsed() {
        Ok(elapsed) => {
            println!("{}", elapsed.as_millis());
        }
        Err(e) => {
            println!("Error: {:?}", e);
        }
    }
    /*
    while let Some(top) = stack.pop() {
        println!("{}", top);
    }
    */
}

pub fn merge_sort(mut input: Vec<usize>) -> Vec<usize> {
    let length = input.len();
    // Base case
    if length < 2 {
        return input;
    }

    let right = merge_sort(input.split_off(length / 2));
    let left = merge_sort(input);

    let mut left_iter = left.iter().peekable();
    let mut right_iter = right.iter().peekable();

    let mut merged = Vec::with_capacity(length);
    while let (Some(&x), Some(&y)) = (left_iter.peek(), right_iter.peek()) {
        if *x <= *y {
            merged.push(*(left_iter.next().unwrap()))
        } else {
            merged.push(*(right_iter.next().unwrap()))
        }
    }

    for x in left_iter {
        merged.push(*x)
    }

    for y in right_iter {
        merged.push(*y)
    }

    merged
}