
use super::Day;
use alex_lib::read_lines;

#[derive(Debug)]
pub struct Day04 {
    input: Vec<String>,
}

impl Day04 {
    pub fn new() -> Day04 {
        Day04 { input: Vec::new() }
    }

    fn parse_input(&mut self) {}
}

impl Day for Day04 {
    fn day_nr(&self) -> String {
        String::from("04")
    }
    fn title(&self) -> String {
        String::from("xxxxxxxxxxxx")
    }

    fn prepare(&mut self) {
        let input = read_lines("data/day04.txt");
        // let input = read_lines("data/day04-test.txt");
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
