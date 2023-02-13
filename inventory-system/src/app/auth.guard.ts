import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import {
  ActivatedRouteSnapshot,
  CanActivate,
  Router,
} from "@angular/router";

@Injectable({
  providedIn: "root",
})
export class AuthGuard implements CanActivate {
  constructor(
    private router: Router,
    private httpClient: HttpClient
  ) {}

  canActivate(route: ActivatedRouteSnapshot) {
    this.httpClient.get('http://localhost:8080/user-session')
    .subscribe((res) => {
        console.log(res, 'AuthGuard')
    })
    if(localStorage.getItem('token')){
        return true
    }else{
        return false
    }
    
  }

}