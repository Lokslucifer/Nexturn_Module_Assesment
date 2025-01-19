describe("Check array accessing", function () {
    it("checking empty array", function () {
        var arr=[]
      expect(getElement(arr,0)).toBe("Index out of bounds");
      expect(getElement(arr,1)).toBe("Index out of bounds");
    });

    it("check non empty array", function () {
        var arr=[1,2,3]
        expect(getElement(arr,-1)).toBe("Index out of bounds");
        expect(getElement(arr,100)).toBe("Index out of bounds");
        expect(getElement(arr,1)).toBe(2);
        expect(getElement(arr,0)).toBe(1);
    });

  });