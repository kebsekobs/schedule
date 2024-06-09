import React, { useEffect, useState } from "react";
import { useForm, Controller } from "react-hook-form";
import styles from "../../shared/style/modal.module.css";
import Button from "../../../components/button";
import { useEditDisciplinesMutation } from "../api/editDisciplinesMutation";
import { useDisciplinesByIdQuery } from "../api/geDisciplinesById";
import { useGroupsQuery } from "../../groups/api/getGroupsQuery";
import { useTeachersQuery } from "../../teachers/api/getTeachersQuery";
import { Multiselect } from "multiselect-react-dropdown";

const EditDisciplinesModal = ({ isOpen, toggleModal, id }) => {
  const disciplinesByIdQuery = useDisciplinesByIdQuery(id);
  const [isDataLoaded, setIsDataLoaded] = useState(false);

  const {
    register,
    control,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm({
    defaultValues: {
      classroomId: "",
      capacity: "",
    },
  });

  const editDisciplinesMutation = useEditDisciplinesMutation();
  const getGroup = useGroupsQuery().data;
  const teachersQuery = useTeachersQuery();

  const options = getGroup?.map((el) => `${el.groupId} ${el.id}`);

  const onSubmit = (data) => {
    data.id = id;
    editDisciplinesMutation.mutateAsync(data);
    toggleModal();
    reset();
  };

  useEffect(() => {
    if (disciplinesByIdQuery.data) {
      reset({
        disciplinesId: disciplinesByIdQuery.data.disciplinesId,
        name: disciplinesByIdQuery.data.name,
        teachers: disciplinesByIdQuery.data.teachers,
        hours: disciplinesByIdQuery.data.hours,
        relatedGroupsId: disciplinesByIdQuery.data.relatedGroupsId,
      });
      setIsDataLoaded(true);
    }
  }, [disciplinesByIdQuery.data, reset]);

  if (!isOpen) {
    return null;
  }

  if (!isDataLoaded) {
    return <div>Loading...</div>;
  }

  const multiSelectStyles = {
    multiselectContainer: {
      backgroundColor: "white",
    },
    inputField: {
      margin: "5px",
    },
    chips: {
      background: "#8090bc",
    },
    optionContainer: {
      maxHeight: "160px",
      overflowY: "scroll",
    },
  };

  return (
    <div className={styles["backdrop"]}>
      <div className={styles["modal"]}>
        <div className={styles["modal-header"]}>
          <h2>Измените данные</h2>
          <Button onClick={toggleModal} styleFeature="close">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              x="0px"
              y="0px"
              width="20"
              height="20"
              viewBox="0 0 30 30"
            >
              <path d="M 7 4 C 6.744125 4 6.4879687 4.0974687 6.2929688 4.2929688 L 4.2929688 6.2929688 C 3.9019687 6.6839688 3.9019687 7.3170313 4.2929688 7.7070312 L 11.585938 15 L 4.2929688 22.292969 C 3.9019687 22.683969 3.9019687 23.317031 4.2929688 23.707031 L 6.2929688 25.707031 C 6.6839688 26.098031 7.3170313 26.098031 7.7070312 25.707031 L 15 18.414062 L 22.292969 25.707031 C 22.682969 26.098031 23.317031 26.098031 23.707031 25.707031 L 25.707031 23.707031 C 26.098031 23.316031 26.098031 22.682969 25.707031 22.292969 L 18.414062 15 L 25.707031 7.7070312 C 26.098031 7.3170312 26.098031 6.6829688 25.707031 6.2929688 L 23.707031 4.2929688 C 23.316031 3.9019687 22.682969 3.9019687 22.292969 4.2929688 L 15 11.585938 L 7.7070312 4.2929688 C 7.5115312 4.0974687 7.255875 4 7 4 z"></path>
            </svg>
          </Button>
        </div>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className={styles["groups-form"]}
        >
          <label>Измените id дисциплины</label>
          <input
            {...register("disciplinesId", { required: true })}
            placeholder="A41E"
            className={styles["input"]}
          />
          {errors.id && (
            <span className={styles.error}>Это поле обязательно</span>
          )}
          <label>Измените имя дисциплины</label>
          <input
            {...register("name", { required: true })}
            placeholder="Ин.яз"
            className={styles["input"]}
          />
          <label>Измените кол-во часов в неделю</label>
          <input
            {...register("hours", { required: true })}
            placeholder="20"
            className={styles["input"]}
          />
          <label>Измените преподователя</label>
          <select
            {...register("teachers", { required: true })}
            className={styles["input"]}
          >
            {teachersQuery.data.map((el, index) => (
              <option
                key={index}
                value={`${el.id}`}
              >{`Преподователь: ${el.name}`}</option>
            ))}
          </select>
          <label>Измените группу(ы)</label>
          <Controller
            control={control}
            name="relatedGroupsId"
            render={({ field: { value, onChange } }) => (
              <Multiselect
                placeholder="314 GTH-JDO-NSK"
                options={options}
                isObject={false}
                showCheckbox={true}
                hidePlaceholder={true}
                closeOnSelect={false}
                onSelect={onChange}
                onRemove={onChange}
                selectedValues={value}
                showArrow={true}
                style={multiSelectStyles}
              />
            )}
          />
          <Button type="submit">Oтправить</Button>
        </form>
      </div>
    </div>
  );
};

export default EditDisciplinesModal;
