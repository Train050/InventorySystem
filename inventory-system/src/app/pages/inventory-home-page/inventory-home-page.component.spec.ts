import { ComponentFixture, TestBed, async, fakeAsync, tick } from '@angular/core/testing';
import { InventoryHomePageComponent } from './inventory-home-page.component';
import { FormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatInputModule } from '@angular/material/input';
import { MatPaginatorModule, MatPaginator } from '@angular/material/paginator';
import { MatSortModule } from '@angular/material/sort';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NO_ERRORS_SCHEMA } from '@angular/core';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { fn } from 'cypress/types/jquery';
import { HttpClient, HttpHandler } from '@angular/common/http';
import { HarnessLoader } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { MatButtonHarness } from '@angular/material/button/testing';
import { Dialog } from '@angular/cdk/dialog';
import { AddItemComponent } from 'app/pages/inventory-home-page/add-item';
import { HttpClientTestingModule } from '@angular/common/http/testing';

fdescribe('InventoryHomePageComponent', () => {
  let component: InventoryHomePageComponent;
  let fixture: ComponentFixture<InventoryHomePageComponent>;
  let loader: HarnessLoader;
  let dialog: MatDialog;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      providers: [{provide: MatDialog, useValue: {open: () => {}}}],
      imports: [HttpClientTestingModule, MatFormFieldModule, MatGridListModule, MatPaginatorModule, FormsModule, MatInputModule, MatSortModule, BrowserAnimationsModule, MatDialogModule],
      declarations: [ InventoryHomePageComponent ],
      schemas: [NO_ERRORS_SCHEMA]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InventoryHomePageComponent);
    loader = TestbedHarnessEnvironment.loader(fixture);
    component = fixture.componentInstance;
    dialog = TestBed.inject(MatDialog);
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
    // expect(component.dataSource).toBeTruthy();
    // expect(component.paginator).toBeInstanceOf(MatPaginator);
    // expect(component.applyFilter).toBeTruthy();
    // expect(component).toBeDefined();
    // expect(component.openAddItem).toBeTruthy();
  });
  it('should set the filter value', () => {
    component.applyFilter(<any>{target: {value: 'test'}});
    expect(component.dataSource.filter).toBe('test');
  });
  it('should open the add item dialog', async () => {
    spyOn(dialog, 'open');
    const button = await loader.getHarness(MatButtonHarness.with({text: 'Add Item'}));
    await button.click();
    expect(dialog.open).toHaveBeenCalledWith(AddItemComponent);
  });
});


