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
  registerUserName: string | null = null
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
    //console.log(postData);
    const postData = {
      
      userName: this.registerUserName,
      email: this.registerEmail,
      password: this.registerPassword,
      phoneNumber: this.registerPhoneNumber,
    };
    console.log(postData);
    this.httpClient.post('http://localhost:8080/registration', postData)
      .subscribe((response: any) => {
      console.log(response);
      if(response){
        localStorage.setItem('token', response.jwt)
        this.router.navigate(['inventory-home-page'])
      }
      this.registerUserName = null
      this.registerEmail = null
      this.registerPassword = null
      this.registerPhoneNumber = null
    })
  }
  login(){
    console.log("in login")
    const postData = {
      email: String,
      password: String
    };
    console.log(postData);
    this.httpClient.get('http://localhost:8080/login') 
      .subscribe((response: any) => {
      console.log(response);
      if(response){
        localStorage.setItem('token', response.jwt)
        this.router.navigate(['inventory-home-page'])
      }
      this.loginEmail = null
      this.loginPassword = null
    })
  }
}