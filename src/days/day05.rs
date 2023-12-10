use super::Day;
use alex_lib::read_lines;

#[derive(Debug)]
pub struct Day05 {
    input: Vec<String>,
}

impl Day05 {
    pub fn new() -> Day05 {
        Day05 { input: Vec::new() }
    }

    fn parse_input(&mut self) {}
}

impl Day for Day05 {
    fn day_nr(&self) -> String {
        String::from("05")
    }
    fn title(&self) -> String {
        String::from("xxxxxxxxx")
    }

    fn prepare(&mut self) {
        let input = read_lines("data/day05.txt");
        // let input = read_lines("data/day05-test.txt");
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
