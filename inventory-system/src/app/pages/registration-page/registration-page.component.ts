import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './registration-page.component.html',
  styleUrls: ['./registration-page.component.scss']
})
export class RegistrationPageComponent {
  username = "";
  password = "";
  email = "";
  phoneNumber = "";
  errorMsg = "";

  constructor(private auth: AuthService, private router: Router) { }
  ngOnInit(): void {
  }

  registration() {
    if (this.username.trim().length === 0) {
        this.errorMsg = "Username is required";
    } 
    else if (this.password.trim().length === 0) {
        this.errorMsg = "Password is required";
    } 
    else if (this.email.trim().length === 0)
    {
        this.errorMsg = "Email is required";
    }
    else if (this.phoneNumber.trim().length === 0)
    {
        this.errorMsg = "Phone number is required"
    }
    else {
      this.errorMsg = "";
      let res = this.auth.registration(this.username, this.password, this.email, this.phoneNumber);
      if (res === 200) {
        this.router.navigate(['home'])
      }
      if (res === 403) {
        this.errorMsg = "Invalid Credentials";
      }
    }
  }
}
