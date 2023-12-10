use alex_lib::read_lines;

use super::Day;

#[derive(Debug)]
pub struct Day02 {
    input: Vec<String>,
}

impl Day02 {
    pub fn new() -> Day02 {
        Day02 { input: Vec::new() }
    }

    fn parse_input(&mut self) {}
}

impl Day for Day02 {
    fn day_nr(&self) -> String {
        String::from("02")
    }
    fn title(&self) -> String {
        String::from("xxxx")
    }

    fn prepare(&mut self) {
        let input = read_lines("data/day02.txt");
        // let input = read_lines("data/day02-test.txt");
        self.input = input;
        self.parse_input();
    }

    fn solve1(&mut self) -> String {
        let mut sum = 0;
        String::from(format!("{0}", sum))
    }
    fn solve2(&mut self) -> String {
        let mut sum = 0;
        String::from(format!("{0}", sum))
    }
}
