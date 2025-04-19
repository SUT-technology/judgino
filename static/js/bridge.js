document.addEventListener("DOMContentLoaded", function () {
    const select = document.getElementById("role-select");
    select.addEventListener("change", function () {
        const selectedRole = select.value;
        fetch(`/profile/{{.UserId}}/change_role`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ role: selectedRole })
        })
        .then(response => {
            if (response.ok) {
                alert("Role changed successfully!");
                window.location.reload(); // optional
            } else {
                alert("Failed to change role.");
            }
        });
    });
});