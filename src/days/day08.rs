use super::Day;
use alex_lib::read_lines;

#[derive(Debug)]
pub struct Day08 {
    input: Vec<String>,
}

impl Day08 {
    pub fn new() -> Day08 {
        Day08 { input: Vec::new() }
    }

    fn parse_input(&mut self) {}
}

impl Day for Day08 {
    fn day_nr(&self) -> String {
        String::from("08")
    }
    fn title(&self) -> String {
        String::from("xxxxxxx")
    }

    fn prepare(&mut self) {
        let input = read_lines("data/day08.txt");
        // let input = read_lines("data/day08-test.txt");
        self.input = input;
        self.parse_input();
    }

    fn solve1(&mut self) -> String {
        let mut solution: u64 = 0;
        String::from(format!("{0}", solution))
    }

    fn solve2(&mut self) -> String {
        let mut solution: u64 = 0;
        String::from(format!("{0}", solution))
    }
}
