---
title: "Register"
layout: "register.html"

---

    <h2>Register</h2>
    <form action="/register" method="post">
        <div>
            <label for="username">Username</label>
            <input type="text" id="username" name="username" required>
        </div>
        <div>
            <label for="password">Password</label>
            <input type="password" id="password" name="password" required>
        </div>
        <div>
            <input type="submit" value="Register">
        </div>
    </form>
    <p>Already have an account? <a href="/login">Login</a></p>
