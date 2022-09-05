struct Student {
    char name[20];
    int  *age[20];
    int  *year;
    float gpa;
};

 struct Organisation
{
  char organisation_name[20];
  char org_number[20];

  // Dependent structure is used
  // as a member inside the main
  // structure for implementing
  // nested structure
  struct Student emp;
};

struct mabbas {
  bool mnum;
  char mlet;
};