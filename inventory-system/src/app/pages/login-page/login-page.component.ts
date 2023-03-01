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
  registerPassword: string | null = null
  loginEmail: string | null = null
  loginPassword: string | null = null
  constructor(
    private httpClient: HttpClient,
    private router: Router
  ){
  }
  register(){
    console.log(this.registerFirstName, this.registerPassword)
    this.httpClient.post('http://localhost:8080/register', {
      firstName: this.registerFirstName,
      lastName: this.registerLastName,
      email: this.registerEmail,
      password: this.registerPassword,
    }).subscribe((response: any) => {
      if(response){
        localStorage.setItem('token', response.jwt)
        this.router.navigate(['profile'])
      }
      this.registerFirstName = null
      this.registerLastName = null
      this.registerEmail = null
      this.registerPassword = null
    })
  }
  login(){
    this.httpClient.post('http://localhost:8080/login', {
      email: this.loginEmail,
      password: this.loginPassword
    }).subscribe((response: any) => {
      if(response){
        localStorage.setItem('token', response.jwt)
        this.router.navigate(['profile'])
      }
      this.loginEmail = null
      this.loginPassword = null
    })
  }
}