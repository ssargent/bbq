using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using BbqStore.Core.Entities;
using BbqStore.Core.Services;
using Marten;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace BbqStore.WebApp.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class StoreController : ControllerBase
    {
        protected IDocumentSession DocumentSession { get; set; }
     
        public StoreController(IDocumentSession documentSession )
        {
            DocumentSession = documentSession; 
        }

        [HttpGet, Route("")]
        public async Task<IActionResult> Get()
        {
            return Ok(DocumentSession.Query<Store>().Where(s => s.IsDeleted == false));
        }

        [HttpGet, Route("{key}")]
        public async Task<IActionResult> Get(string key)
        {
            return Ok(DocumentSession.Query<Store>().Where(s => s.IsDeleted == false && s.Key == key));
        }
    }

 
}