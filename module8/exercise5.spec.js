describe("DOM Manipulation: toggleVisibility", function () {
    let element;
  
    beforeEach(function () {
      element = {
        style: {
          display: "block",
        },
      };
      spyOnProperty(element.style, "display", "get").and.callThrough();
      spyOnProperty(element.style, "display", "set").and.callThrough();
    });
  
    it("should hide the element if it is initially visible", function () {
      toggleVisibility(element);
      expect(element.style.display).toBe("none");
      expect(element.style.display).toHaveBeenCalledTimes(2); 
    });
  
    it("should show the element if it is initially hidden", function () {
      element.style.display = "none";
      toggleVisibility(element);
      expect(element.style.display).toBe("block");
      expect(element.style.display).toHaveBeenCalledTimes(2);
    });
  });
  