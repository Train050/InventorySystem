import { Component, ViewChild } from '@angular/core';
import {MatPaginator} from '@angular/material/paginator';
import {MatSort} from '@angular/material/sort';
import {MatTableDataSource} from '@angular/material/table';
import {MatDialog} from '@angular/material/dialog';
import { AddItemComponent } from 'app/pages/inventory-home-page/add-item';
import { HttpClient } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { Router } from '@angular/router';

export interface UserData {
  ID: number;
  ProductName: string;
  DateAcquired: string;
  ProductAmount: number;
}

/** Constants used to fill up our data base. */
const FRUITS: string[] = [
  'blueberry',
  'lychee',
  'kiwi',
  'mango',
  'peach',
  'lime',
  'pomegranate',
  'pineapple',
];
const NAMES: string[] = [
  'Maia',
  'Asher',
  'Olivia',
  'Atticus',
  'Amelia',
  'Jack',
  'Charlotte',
  'Theodore',
  'Isla',
  'Oliver',
  'Isabella',
  'Jasper',
  'Cora',
  'Levi',
  'Violet',
  'Arthur',
  'Mia',
  'Thomas',
  'Elizabeth',
];

@Component({
  selector: 'app-inventory-home-page',
  templateUrl: './inventory-home-page.component.html',
  styleUrls: ['./inventory-home-page.component.scss']
})
export class InventoryHomePageComponent {
  displayedColumns: string[] = ['ID', 'ProductName', 'DateAcquired', 'ProductAmount'];
  dataSource: MatTableDataSource<UserData>;

  @ViewChild(MatPaginator) paginator:any = MatPaginator;
  @ViewChild(MatSort) sort:any = MatSort;

 public readonly inventoryItems$: Observable<UserData[]> = this.httpClient.get<UserData[]>('http://localhost:4200/api/inventory');

 constructor(public dialog: MatDialog, private readonly httpClient: HttpClient, private router: Router) {
    
    // Create 100 users
    const users = Array.from({length: 100}, (_, k) => createNewUser(k + 1));

    // Assign the data to the data source for the table to render
    this.dataSource = new MatTableDataSource();
    this.inventoryItems$.subscribe((items) => { this.dataSource.data = items; this.dataSource.paginator = this.paginator;
      this.dataSource.sort = this.sort;});
  }

  ngAfterViewInit() {
    // this.dataSource.paginator = this.paginator;
    // this.dataSource.sort = this.sort;
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();

    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage();
    }
  }

  openAddItem() {
    const dialogRef = this.dialog.open(AddItemComponent)
  }

  logout() {
    this.router.navigate(['/login-page']);
  }
}

/** Builds and returns a new User. */
function createNewUser(id: number): UserData {
  const name =
    NAMES[Math.round(Math.random() * (NAMES.length - 1))] +
    ' ' +
    NAMES[Math.round(Math.random() * (NAMES.length - 1))].charAt(0) +
    '.';

  return {
    ID: id,
    ProductName: name,
    DateAcquired: new Date().toDateString(),
    ProductAmount: parseInt(FRUITS[Math.round(Math.random() * (FRUITS.length - 1))]),
  };
}
