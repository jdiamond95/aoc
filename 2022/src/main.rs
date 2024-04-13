use std::env::args;
use std::fs;

fn main() {
    let args: &Vec<String> = &args().collect();

    let days: Vec<u8> = match args.len() {
        2 => vec![args[1].parse().unwrap()],
        _ => (1..=3).collect(),
    };

    for day in days {
        let path = format!("./src/inputs/day{:02}.txt", day);
        let input = fs::read_to_string(&path).unwrap();

        match day {
            1 => aoc2022::day01::run(&input),
            2 => aoc2022::day02::run(&input),
            3 => aoc2022::day03::run(&input),
            _ => println!("Sorry havent done that one yet"),
        }
    }
}
