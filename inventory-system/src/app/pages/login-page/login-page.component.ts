import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../services/auth.service';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.scss']
})

export class LoginPageComponent {
  
  isLogin: boolean = false
  registerFirstName: string | null = null
  registerLastName: string | null = null
  registerEmail: string | null = null
  registerPhoneNumber: string | null = null
  registerPassword: string | null = null
  loginEmail: string | null = null
  loginPassword: string | null = null
  //http://localhost:8080/login

  constructor(
    private httpClient: HttpClient,
    private router: Router
  ){
  }
  register(){
    console.log("in register");
    const postData = {
      
      firstName: this.registerFirstName,
      lastName: this.registerLastName,
      email: this.registerEmail,
      password: this.registerPassword,
      phoneNumber: this.registerPhoneNumber,
    };
    console.log(postData);
    this.httpClient.post('http://localhost:8080/register', postData)
    .subscribe((response: any) => {
      console.log(response);
      if(response){
        localStorage.setItem('token', response.jwt)
        this.router.navigate(['profile'])
      }
      this.registerFirstName = null
      this.registerLastName = null
      this.registerEmail = null
      this.registerPassword = null
      this.registerPhoneNumber = null
    })
  }
  login(){
    console.log("in login")
    const postData = {
      email: this.loginEmail,
      password: this.loginPassword,
    };
    console.log(postData);
    this.httpClient.post('http://localhost:8080/login', postData) 
      .subscribe((response: any) => {
      if(response){
        localStorage.setItem('token', response.jwt)
        this.router.navigate(['profile'])
      }
      this.loginEmail = null
      this.loginPassword = null
    })
  }
}