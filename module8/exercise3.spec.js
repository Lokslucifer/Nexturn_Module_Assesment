describe("Async Functions: Delayed Greeting", function () {
    beforeEach(function () {
      jasmine.clock().install();
    });
  
    afterEach(function () {
      jasmine.clock().uninstall();
    });
  
    it("should resolve with the correct greeting message", async function (done) {
      const promise = delayedGreeting("John", 1000);
      jasmine.clock().tick(1000);
      promise.then((message) => {
        expect(message).toBe("Hello, John!");
        done();
      });
    });
  
    it("should respect the specified delay", function () {
      const delay = 2000;
      const callback = jasmine.createSpy("callback");
  
      delayedGreeting("Jane", delay).then(callback);
      jasmine.clock().tick(delay - 1);
      expect(callback).not.toHaveBeenCalled();
  
      jasmine.clock().tick(1);
      expect(callback).toHaveBeenCalled();
    });
  });
  