describe("String Utilities", function () {
    describe("capitalize", function () {
      it("should capitalize the first letter of a word", function () {
        expect(capitalize("hello")).toBe("Hello");
        expect(capitalize("world")).toBe("World");
      });
  
      it("should return an empty string if the input is an empty string", function () {
        expect(capitalize("")).toBe("");
      });
  
  
      it("should not change the case of the rest of the string", function () {
        expect(capitalize("hELLO")).toBe("HELLO");
        expect(capitalize("wORLD")).toBe("WORLD");
      });
  
      it("should return an empty string if the input is undefined or null", function () {
        expect(capitalize(undefined)).toBe("");
        expect(capitalize(null)).toBe("");
      });
    });
  
    describe("reverseString", function () {
      it("should reverse a string", function () {
        expect(reverseString("hello")).toBe("olleh");
        expect(reverseString("world")).toBe("dlrow");
      });
  
      it("should return an empty string if the input is an empty string", function () {
        expect(reverseString("")).toBe("");
      });
  
      it("should handle palindromes correctly", function () {
        expect(reverseString("madam")).toBe("madam");
        expect(reverseString("racecar")).toBe("racecar");
      });

  
      it("should reverse strings with special characters", function () {
        expect(reverseString("123!@#")).toBe("#@!321");
        expect(reverseString("hello, world!")).toBe("!dlrow ,olleh");
      });
    });
  });
  