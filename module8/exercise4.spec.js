describe("Notification Service", function () {
    let notificationService;
  
    beforeEach(function () {
      notificationService = {
        send: jasmine.createSpy("send"),
      };
    });
  
    it("should return 'Notification Sent' when the notification is successfully sent", function () {
      notificationService.send.and.returnValue(true);
      const result = sendNotification(notificationService, "Hello, World!");
      expect(notificationService.send).toHaveBeenCalledWith("Hello, World!");
      expect(result).toBe("Notification Sent");
    });
  
    it("should return 'Failed to Send' when the notification fails to send", function () {
      notificationService.send.and.returnValue(false);
      const result = sendNotification(notificationService, "Hello, World!");
      expect(notificationService.send).toHaveBeenCalledWith("Hello, World!");
      expect(result).toBe("Failed to Send");
    });
  });
  