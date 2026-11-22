$(document).ready(function () {
    $("form").on("submit", function (event) {
        event.preventDefault();

        const form = $(this);
        const password = $("#password").val();
        const confirmPassword = $("#confirmPassword").val();

        if (password !== confirmPassword) {
            alert("Passwords do not match!");
            return;
        }

        const btn = $("#signup-form-button");
        btn.text("Sending...");

        setTimeout(function () {
            $.ajax({
                url: "/users",
                method: "POST",
                data: form.serialize(),
                success: function (data) {
                    alert("You have successfully registered!");
                },
                error: function (err) {
                    console.log(err);
                    if (err?.responseJSON?.error) {
                        alert(err.responseJSON.error);
                    }else{
                        alert("Error in registration.");
                    }
                },
                complete: function () {
                    btn.text("Create Account");
                }
            });
        }, 200)

    });
});