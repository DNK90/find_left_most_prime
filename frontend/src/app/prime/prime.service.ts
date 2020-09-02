import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';

import { Prime } from './prime';
import { HttpErrorHandler, HandleError } from '../http-error-handler.service';
import { environment } from '../../environments/environment';

const httpOptions = {
	headers: new HttpHeaders({
		'Content-Type':  'application/json'
	})
}

@Injectable()
export class PrimeService {
	apiUrl = new String(environment.apiUrl).concat("/v1/prime");
	private handleError: HandleError;

	constructor(private http:HttpClient, httpErrorHandler: HttpErrorHandler) {
		this.handleError = httpErrorHandler.createHandleError("PrimeService");
	}

	getPrime(no: string): Observable<Prime> {
		console.log(this.apiUrl);
		const options = {
			params: new HttpParams().set("number", no)
		};
		return this.http.get<Prime>(this.apiUrl, options).pipe(
			catchError(this.handleError<Prime>("getPrime", { output: 0, elapsedTime: 0 }))
		);
	}
}