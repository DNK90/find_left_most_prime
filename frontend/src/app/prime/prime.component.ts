import { Component, OnInit } from '@angular/core';
import { PrimeService } from './prime.service';
import { Prime } from './prime';

@Component({
  selector: 'app-prime',
  providers: [PrimeService],
  templateUrl: 'prime.component.html',
  styleUrls: ['./prime.component.css']
})
export class PrimeComponent implements OnInit {
  prime: Prime;
  constructor(private primeService: PrimeService) { }
  ngOnInit(): void {
  }
  getPrime(no: string): void {
    this.primeService.getPrime(no).subscribe(prime => (this.prime = prime));
  }

}
