import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.scss']
})
export class LoginPageComponent {
  username = "";
  password = "";
  errorMsg = "";

  constructor(private auth: AuthService, private router: Router) { }
  ngOnInit(): void {
  }

  login() {
    if (this.username.trim().length === 0) {
        this.errorMsg = "Username is required";
    } else if (this.password.trim().length === 0) {
        this.errorMsg = "Password is required";
    } else {
      this.errorMsg = "";
      let res = this.auth.login(this.username, this.password);
      if (res === 200) {
        this.router.navigate(['home'])
      }
      if (res === 403) {
        this.errorMsg = "Invalid Credentials";
      }
    }
  }
}
